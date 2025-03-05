import { useLocation } from "preact-iso";
import { navbarroutes } from "./routes";
import { AllDaisyUiThemes } from "./themes/AllDaisyUiThemes";

interface SiteLayoutProps {
  children: React.ReactNode;
}

export function SiteLayout({ children }: SiteLayoutProps) {
  const { url } = useLocation();
  const closeDrawer = () => {
    const drawer =
      document.getElementById("site-layout-side-drawer") as HTMLInputElement;
      console.log(drawer);
    drawer?.click();
  }
  return (
    <div className="w-full h-full flex flex-col items-center justify-center">
      <div className="drawer">
        <input id="site-layout-side-drawer" type="checkbox" className="drawer-toggle" />
        <div className="drawer-content flex flex-col">
          {/* Navbar */}
          <div className="navbar bg-base-300 w-full sticky top-0 z-30">
            <div className="flex-none lg:hidden">
              <label
                htmlFor="site-layout-side-drawer"
                aria-label="open sidebar"
                className="btn btn-square btn-ghost">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  class="inline-block h-6 w-6 stroke-current">
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    d="M4 6h16M4 12h16M4 18h16"></path>
                </svg>
              </label>
            </div>
            <div className="mx-2 flex-1 px-2">
              <a href="/" className={`btn btn-ghost `}>
                <div className="text-xl font-bold">Cloud Mwitu</div>
              </a>
            </div>
            <div className="hidden flex-none lg:block">
              <ul className="menu menu-horizontal gap-5 items-center">
                {/* Navbar menu content here */}
                <li className={`flex flex-row  gap-3 `}>
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
                </li>
                <li>
                  <AllDaisyUiThemes />
                </li>
              </ul>
            </div>
          </div>
          {/* Page content here */}
          {children}
        </div>
        <div className="drawer-side">
          <label
            htmlFor="site-layout-side-drawer"
            aria-label="close sidebar"
            className="drawer-overlay"></label>
          <ul className="menu bg-base-200 min-h-full justify-between w-80 p-4">
            {/* Sidebar content here */}
            <li className={`flex items-center justify-center gap-28 w-full `}>
              <div className="mx-2 flex-1 justify-center items-center w-full border-b hover:text-primary hover:border-primary">
                <a href="/" className={` `}>
                  <div className="text-xl font-bold">Cloud Mwitu</div>
                </a>
              </div>
              <div className="flex flex-col items-center justify-center gap-3 w-full">
                {navbarroutes.map((route) => (
                  <a
                    href={route.path}
                    onClick={closeDrawer}
                    className={`btn btn-sm w-full ${
                      url === route.path ? "btn-secondary btn-outline" : "btn-ghost"
                    }`}>
                    {route.name}
                    {/* <route.icon size={20}/> */}
                  </a>
                ))}
              </div>
            </li>
            <li className={` w-full`}>
              <AllDaisyUiThemes />
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
}
