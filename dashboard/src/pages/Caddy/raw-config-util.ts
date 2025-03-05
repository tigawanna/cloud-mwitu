import { ConfigItem } from "./types";

export const generateCaddyConfig = (configData: ConfigItem[], domain: string) => {
  // Handle empty config
  if (!configData || configData.length === 0) {
    return `${domain} {}\n`;
  }

  // Find root item (just the domain)
  const rootItem = configData.find(item => 
    item.path.length === 1 && item.path[0] === domain
  );
  
  // Start with the domain
  let configStr = `${domain} {\n`;
  
  // Handle simple domain-level directives (those with exactly 2 path parts: domain and directive)
  const domainDirectives = configData.filter(item => 
    item.path.length === 2 && 
    item.path[0] === domain
  );
  
  // Simple case: No domain directives
  if (domainDirectives.length === 0) {
    // Handle the case where there's only a domain with values
    if (rootItem && rootItem.value.length > 0) {
      rootItem.value.forEach(val => {
        configStr += `    ${val}\n`;
      });
    }
    configStr += '}\n';
    return configStr;
  }
  
  // Process each domain-level directive
  domainDirectives.forEach(item => {
    const directive = item.path[1];
    
    // Check if this directive has nested blocks (path length > 2)
    const hasNestedBlocks = configData.some(nestedItem => 
      nestedItem.path.length > 2 &&
      nestedItem.path[0] === domain &&
      nestedItem.path[1] === directive
    );
    
    if (hasNestedBlocks) {
      // Directive with nested blocks
      configStr += `    ${directive} {\n`;
      
      // Add values for this directive
      if (item.value.length > 0) {
        item.value.forEach(val => {
          configStr += `        ${val}\n`;
        });
      }
      
      // Add nested blocks
      const nestedItems = configData.filter(nestedItem => 
        nestedItem.path.length > 2 &&
        nestedItem.path[0] === domain &&
        nestedItem.path[1] === directive
      );
      
      // Process nested items by level
      const processedPaths = new Set<string>();
      nestedItems.forEach(nestedItem => {
        const nestedPath = nestedItem.path.slice(2); // Remove domain and parent directive
        const nestedPathStr = nestedPath.join('|');
        
        if (processedPaths.has(nestedPathStr)) {
          return; // Skip if already processed
        }
        
        if (nestedPath.length > 0) {
          configStr += `        ${nestedPath.join(' ')} {\n`;
          nestedItem.value.forEach(val => {
            configStr += `            ${val}\n`;
          });
          configStr += `        }\n`;
          processedPaths.add(nestedPathStr);
        }
      });
      
      configStr += `    }\n`;
    } else {
      // Simple directive with values but no nested blocks
      if (item.value.length > 0) {
        item.value.forEach(val => {
          configStr += `    ${directive} ${val}\n`;
        });
      } else {
        configStr += `    ${directive}\n`;
      }
    }
  });
  
  // Close the domain block
  configStr += '}\n';
  return configStr;
};
