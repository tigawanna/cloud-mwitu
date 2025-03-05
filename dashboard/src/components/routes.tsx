import { CiHome } from "react-icons/ci";
import { TbError404 } from "react-icons/tb";
export const navbarroutes = [
  {
    name: "Home",
    path: "/",
    icon: CiHome,
  },
  {
    name: "Caddy",
    path: "/caddy",
    icon: CiHome,
  },
  {
    name: "404",
    path: "/404",
    icon: TbError404,
  },
] as const
