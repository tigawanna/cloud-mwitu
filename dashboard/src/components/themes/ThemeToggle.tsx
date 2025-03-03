import { ViewTransitionSelect } from "./ViewTransitionSelect";
import { AllDaisyUiThemes } from "./AllDaisyUiThemes";
import { MdOutlineWbSunny } from "react-icons/md";
import { HiOutlineMoon } from "react-icons/hi";
import { useTheme } from "./use-theme";


interface ThemeToggleProps {
  compact?: boolean;
}

export function ThemeToggle({ compact }: ThemeToggleProps) {
  const { theme, updateTheme } = useTheme();

  function transitionColors() {
    if (typeof window !== "undefined") {
      try {
        document.startViewTransition(() => {
          const newTheme = theme === "light" ? "dark" : "light";
          document.documentElement.dataset.theme = newTheme;
          updateTheme(newTheme);
        });
      } catch (error) {
        const newTheme = theme === "light" ? "dark" : "light";
        document.documentElement.dataset.theme = newTheme;
        updateTheme(newTheme);
      }
    }
  }
  return (
    <div className="flex flex-wrap w-full items-center justify-between gap-4 px-2">
      <div className="flex w-full items-center justify-between gap-4">
        <ViewTransitionSelect compact={compact} />
        <button onClick={() => transitionColors()} className="">
          {theme === "light" ? <HiOutlineMoon /> : <MdOutlineWbSunny  />}
        </button>
      </div>
      <AllDaisyUiThemes compact={compact} />
    </div>
  );
}
