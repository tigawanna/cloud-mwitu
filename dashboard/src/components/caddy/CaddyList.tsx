import { useQuery$ } from "@preact-signals/query";
import { Suspense } from "preact/compat";
import { Show, For } from "@preact-signals/utils/components";
import { CardsListSuspenseFallback } from "../shared/CardsListSuspenseFallback";
interface CaddyListProps {

}

export function CaddyList({}:CaddyListProps){
    const query = useQuery$(() => ({
      queryKey: ["caddy"],
      queryFn: fetchStatistics,
      suspense: true,
    }));
return (
  <div className="w-full h-full flex flex-col items-center justify-center">
    <Suspense fallback={<CardsListSuspenseFallback />}>
      <Show when={() => query.data}>
        {(data) => (
          <ul>
            {data().map((item) => (
              <li key={item.label}>{item.data}</li>
            ))}
          </ul>
        )}
      </Show>
    </Suspense>
  </div>
);
}

async function fetchStatistics(){
    const res = await fetch('https://api.caddyserver.com/v2/stats');
    const data = await res.json();
    return data;
}
