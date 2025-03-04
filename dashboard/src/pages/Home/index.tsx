import { Suspense } from "preact/compat";
import preactLogo from "../../assets/preact.svg";
import { Status } from "./Status";
import "./style.css";
import { CaddyList } from "@/components/caddy/CaddyList";
import { CardsListSuspenseFallback } from "@/components/shared/CardsListSuspenseFallback";

export function Home() {
  return (
    <div class="min-h-screen flex flex-col items-center justify-center space-y-8">
      <a href="https://preactjs.com" target="_blank">
        <img src={preactLogo} alt="Preact logo" height="160" width="160" />
      </a>
      <h1 class="text-3xl">Get Started building Vite-powered Preact Apps </h1>
      <Suspense fallback={<div class="text-3xl">Loading...</div>}>
        <Status />
      </Suspense>
          <Suspense fallback={<CardsListSuspenseFallback />}>
      <CaddyList/>
          </Suspense>
    </div>
  );
}
