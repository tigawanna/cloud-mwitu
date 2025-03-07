import { useQuery$ } from "@preact-signals/query";
import { Suspense } from "preact/compat";
import { Show, For } from "@preact-signals/utils/components";
import { CardsListSuspenseFallback } from "../../components/shared/CardsListSuspenseFallback";
import { caddyService } from "@/lib/kubb/gen";
import { RecursiveCaddyConfig } from "./RecursiveCaddyConfig";
interface CaddyListProps {}

export function CaddyList({}: CaddyListProps) {
  const query = useQuery$(() => ({
    queryKey: ["caddy"],
    queryFn: fetchStatistics,
    suspense: true,
  }));
console.log("=== data ==",query.data)
  return (
    <div className="w-full h-full flex flex-col items-center justify-center">
      <Suspense fallback={<CardsListSuspenseFallback />}>
        <Show when={() => (query.data.type === "success" ? query.data.data : [])}>
          {(data) => (
            <ul className="grid p-5 grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-4 overflow-auto">
              {data.map((item) => (
                <li key={item.domain} className="card bg-base-200 m-2">
                  <div className="card-content border-[1px] border-primary/20 rounded-2xl shadow-sm shadow-base-300">
                    <div className="card-body">
                      <RecursiveCaddyConfig configData={item.block || []} domain={item.domain} />
                    </div>
                  </div>
                </li>
              ))}
            </ul>
          )}
        </Show>
      </Suspense>
    </div>
  );
}

async function fetchStatistics() {
  return caddyService().GETcaddyClient();
}
// async function fetchStatistics(){
//   const res2 = await fetch("http://localhost:8080/caddy/");
//   const data = await res2.json()
//   // const res = await fetch("http://localhost:9999/pets/all?per_page=10&page=1")
//   // const data = await res.json()
//   console.log(data)
//   return data

// }
