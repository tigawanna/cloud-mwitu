import { useState } from "react";

interface ViewTransitionSelectProps {
  compact?: boolean
}

export function ViewTransitionSelect({compact}: ViewTransitionSelectProps) {
  const [trns, setTrans] = useState("default");
  
  function changeTransition(value: string) {
    try {
      document.startViewTransition(() => {
        document.documentElement.dataset.style = value;
        setTrans(value);
      });
    } catch (error) {
      document.documentElement.dataset.style = value;
      setTrans(value);
    }
  }

  if (!import.meta.env.DEV || compact) {
    return null;
  }

  const transitions = [
    { value: "default", label: "Default" },
    { value: "vertical", label: "Vertical" },
    { value: "wipe", label: "Wipe" },
    { value: "angled", label: "Angled" },
    { value: "flip", label: "Flip" },
    { value: "slides", label: "Slides" },
  ];

  return (
    <div className="form-control w-full max-w-xs">
      <label className="label">
        <span className="label-text">View Transitions</span>
      </label>
      <select
        value={trns}
        onChange={(e) => changeTransition(e.currentTarget.value)}
        className="select select-bordered select-sm"
      >
        {transitions.map((t) => (
          <option key={t.value} value={t.value}>
            {t.label}
          </option>
        ))}
      </select>
    </div>
  );
}
