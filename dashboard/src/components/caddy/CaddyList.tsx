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
          <ul className="grid p-5 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {data.map((item) => (
              <li key={item.domain} className="card bg-base-200 m-2">
                <div className="card-content">
                  <h3 className="card-title">{item.domain}</h3>
                  <div className="card-body">
                    <p>{item.content}</p>
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

async function fetchStatistics(){
    // const res = await fetch("http://localhost:8080/caddre");
    const res = await fetch("http://localhost:9999/pets/all?per_page=10&page=1");
    const data = await res.json();
    console.log(" caddy list  === ",data);
    return data;
}
