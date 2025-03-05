import { CaddyList } from "@/pages/Caddy/CaddyList";
import { CardsListSuspenseFallback } from "@/components/shared/CardsListSuspenseFallback";
import { Suspense } from "preact/compat";

interface CaddyPageProps {

}

export function CaddyPage({}:CaddyPageProps){
return (
  <div className="w-full h-full flex flex-col items-center justify-center">
    <h1 className="text-4xl font-bold">Caddy Page</h1>
    <Suspense fallback={<CardsListSuspenseFallback />}>
      <CaddyList />
    </Suspense>
  </div>
);
}
