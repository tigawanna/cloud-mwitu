import { useState } from "preact/hooks";
import { ConfigItem, RecursiveConfigProps } from "./types";
import { generateCaddyConfig } from "./raw-config-util";



export function RecursiveCaddyConfig({ configData, domain }: RecursiveConfigProps) {
  // Create a nested structure from the flat config data
  const buildNestedConfig = (items: ConfigItem[]) => {
    const rootNode: any = { children: {}, values: [] };

    items.forEach((item) => {
      let currentNode = rootNode;
      const path = item.path;

      // Navigate through path parts
      for (let i = 0; i < path.length; i++) {
        const pathPart = path[i];
        // Skip the domain as it's the root
        if (i === 0 && pathPart === domain) continue;

        if (!currentNode.children[pathPart]) {
          currentNode.children[pathPart] = { children: {}, values: [] };
        }
        currentNode = currentNode.children[pathPart];
      }

      // Add values to the leaf node
      currentNode.values = item.value;
    });

    return rootNode;
  };

  // Render a config node and its children
  const renderConfigNode = (node: any, pathSoFar: string[] = [], depth = 0) => {
    const [isEditing, setIsEditing] = useState(false);
    const [values, setValues] = useState(node.values);

    const handleValueChange = (index: number, newValue: string) => {
      const newValues = [...values];
      newValues[index] = newValue;
      setValues(newValues);
    };

    const handleAddValue = () => {
      setValues([...values, ""]);
    };

    const handleRemoveValue = (index: number) => {
      const newValues = [...values];
      newValues.splice(index, 1);
      setValues(newValues);
    };

    const handleSave = () => {
      setIsEditing(false);
      // Here you would handle saving changes back to the server
      console.log("Save changes for path:", pathSoFar, "values:", values);
    };

    const childNodes = Object.entries(node.children);
    const hasChildren = childNodes.length > 0;
    const nodeKey = pathSoFar[pathSoFar.length - 1] || domain;

    return (
      <div
        className={`pl-${depth * 4} border-l-2 border-primary ml-2 ${depth > 0 ? "mt-2" : ""}`}
        key={pathSoFar.join("-")}>
        {nodeKey && (
          <div className="flex items-center">
            <span className="font-bold text-md">{nodeKey}</span>
            {values.length > 0 && (
              <button
                className="btn btn-xs btn-ghost ml-2"
                onClick={() => setIsEditing(!isEditing)}>
                {isEditing ? "Cancel" : "Edit"}
              </button>
            )}
          </div>
        )}

        {/* Values section */}
        <div className="pl-4 my-1">
          {values.length > 0 && (
            <>
              {isEditing ? (
                <div className="border border-base-300 rounded-md p-2">
                  {values.map((value, idx) => (
                    <div key={idx} className="flex items-center mb-1">
                      <input
                        type="text"
                        value={value}
                        onChange={(e) =>
                          handleValueChange(idx, (e.target as HTMLInputElement).value)
                        }
                        className="input input-sm input-bordered w-full"
                      />
                      <button
                        className="btn btn-sm btn-error ml-2"
                        onClick={() => handleRemoveValue(idx)}>
                        X
                      </button>
                    </div>
                  ))}
                  <div className="flex gap-2 mt-2">
                    <button className="btn btn-sm btn-success" onClick={handleAddValue}>
                      Add Value
                    </button>
                    <button className="btn btn-sm btn-primary" onClick={handleSave}>
                      Save
                    </button>
                  </div>
                </div>
              ) : (
                <div className="bg-base-300 rounded-md p-2">
                  {values.map((value, idx) => (
                    <div key={idx} className="text-sm font-mono">
                      {value}
                    </div>
                  ))}
                </div>
              )}
            </>
          )}
        </div>

        {/* Children nodes */}
        {hasChildren && (
          <div className="ml-2 mt-1">
            {childNodes.map(([key, childNode]) =>
              renderConfigNode(childNode, [...pathSoFar, key], depth + 1)
            )}
          </div>
        )}
      </div>
    );
  };



  // Get the domain from the first item in the config data
  const effectiveDomain = domain || (configData.length > 0 ? configData[0].path[0] : "");
  const nestedConfig = buildNestedConfig(configData);
  const [showRaw, setShowRaw] = useState(false);
  const rawConfig = generateCaddyConfig(configData, effectiveDomain);

  return (
    <div className="w-full">
      <div className="flex justify-between items-center mb-4">
        <h2 className="text-xl font-bold text-primary">{effectiveDomain} </h2>
        <button className="btn btn-sm btn-outline" onClick={() => setShowRaw(!showRaw)}>
          {showRaw ? "Show Tree View" : "Show Raw Config"}
        </button>
      </div>

      {showRaw ? (
        <div className="bg-base-300 p-4 rounded-lg overflow-auto">
          <pre className="text-sm font-mono whitespace-pre-wrap">{rawConfig}</pre>
        </div>
      ) : (
        <div className="bg-base-100 p-4 rounded-lg shadow-lg">{renderConfigNode(nestedConfig)}</div>
      )}
    </div>
  );
}
