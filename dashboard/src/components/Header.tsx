import { useLocation } from "preact-iso";
import { navbarroutes } from "./rputes";
import { AllDaisyUiThemes } from "./themes/AllDaisyUiThemes";

export function Header() {
  const { url } = useLocation();
  return (
    <div className="navbar bg-base-200 shadow-lg">
      <div className="navbar-start">
        <a href="/" className={`btn btn-ghost `}>
          <div className="text-xl font-bold px-4">Cloud Mwitu</div>
        </a>
      </div>
      <div className="navbar-center">
        <nav className="flex gap-4">
        {navbarroutes.map((route) => (
			<a
            href={route.path}
            className={`btn btn-sm ${
				url === route.path ? "btn-secondary btn-outline" : "btn-ghost"
            }`}>
            {route.name}
			{/* <route.icon size={20}/> */}
          </a>
        ))}
		</nav>
      </div>
      <div className="navbar-end">
		<AllDaisyUiThemes/>
	  </div>
    </div>
  );
}
