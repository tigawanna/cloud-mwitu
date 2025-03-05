import { render } from "preact";
import { LocationProvider, Router, Route, } from "preact-iso";

import { Header } from "./components/Header.jsx";
import { Home } from "./pages/Home/index.jsx";
import { NotFound } from "./pages/_404.jsx";
import "./style.css";
import { MutationCache, QueryClient, QueryClientProvider } from "@preact-signals/query";
import { SiteLayout } from "./components/SiteLayout.js";
import { useEffect } from "preact/hooks";
import { CaddyPage } from "./pages/Caddy/CaddyPage.js";


export const queryClient = new QueryClient({
  mutationCache: new MutationCache({
    onSuccess: async (_, __, ___, mutation) => {
      if (Array.isArray(mutation.meta?.invalidates)) {
        // biome-ignore lint/complexity/noForEach: <explanation>
        mutation.meta?.invalidates.forEach((key) => {
          return queryClient.invalidateQueries({
            queryKey: [key.trim()],
          });
        });
      }
    },
  }),
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 60,
      refetchOnWindowFocus: false,
      refetchOnReconnect: false,
    },
  },
});

export function App() {
    useEffect(() => {
      document.documentElement.dataset.style = "vertical";
    }, []);
  return (
    <QueryClientProvider client={queryClient}>
      <LocationProvider>
        <SiteLayout>
          <main>
            <Router>
              <Route path="/" component={Home} />
              <Route path="/caddy" component={CaddyPage} />
              <Route default component={NotFound} />
            </Router>
          </main>
        </SiteLayout>
      </LocationProvider>
    </QueryClientProvider>
  );
}

render(<App />, document.getElementById("app"));
