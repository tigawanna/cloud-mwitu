import { CaddyList } from "@/pages/Caddy/CaddyList";
import { CardsListSuspenseFallback } from "@/components/shared/CardsListSuspenseFallback";
import { Suspense } from "preact/compat";

interface CaddyPageProps {

}

export function CaddyPage({}:CaddyPageProps){
return (
  <div className="w-full h-full flex flex-col items-center justify-center">
    <Suspense fallback={<CardsListSuspenseFallback />}>
      <CaddyList />
    </Suspense>
  </div>
);
}
