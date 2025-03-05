import { useState } from "react";
interface AllDaisyUiThemesProps {
  compact?: boolean
}

export function AllDaisyUiThemes({compact}: AllDaisyUiThemesProps) {
  const allDaisyUiThems = [
    "forest",
    "cupcake",
    "emerald",
    "corporate",
    "synthwave",
    "retro",
    "luxury",
    "dracula",
    "night",
    "cmyk",
    "autumn",
    "winter",
    "aqua",
    "lofi",
    "pastel",
    "fantasy",
    "wireframe",
    "black",
  ];
  const [theme, setTheme] = useState(allDaisyUiThems[0]);
  
  function changeTheme(theme: string) {
    try {
      document.startViewTransition(() => {
        document.documentElement.dataset.theme = theme;
        setTheme(theme);
      });
    } catch (error) {
      document.documentElement.dataset.theme = theme;
      setTheme(theme);
    }
  }

  if(compact) {
    return null;
  }

  return (
    <select 
      value={theme}
      onChange={(e) => changeTheme(e.currentTarget.value)}
      className="select select-bordered "
    >
      {allDaisyUiThems.map((thm) => (
        <option 
          key={thm} 
          value={thm}
          className={thm === theme ? 'bg-primary text-primary-content' : ''}
        >
          {thm}
        </option>
      ))}
    </select>
  );
}
