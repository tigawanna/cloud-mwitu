import client from '@/lib/kubb/custom-fetch-client.ts'
import type { GETCaddyQueryResponse, GETCaddyQueryParams, GETCaddyHeaderParams, GETCaddy400, GETCaddy500 } from "../../types/'caddyController/GETcaddy.ts"
import type { RequestConfig, ResponseErrorConfig } from '@/lib/kubb/custom-fetch-client.ts'

export function getGETcaddyUrlClient() {
  return `/caddy/` as const
}

/**
 * @description #### Controller: `github.com/tigawanna/cloud-mwitu/internal/controllers.CaddyFileResources.getCaddyFileServices`#### Middlewares:- `github.com/go-fuego/fuego.defaultLogger.middleware`---List all caddyfile services and filter by name
 * @summary get caddy file services
 * {@link /caddy/}
 */
export async function GETcaddyClient(
  params?: GETCaddyQueryParams,
  headers?: GETCaddyHeaderParams,
  config: Partial<RequestConfig> & { client?: typeof client } = {},
) {
  const { client: request = client, ...requestConfig } = config

  const res = await request<GETCaddyQueryResponse, ResponseErrorConfig<GETCaddy400 | GETCaddy500>, unknown>({
    method: 'GET',
    url: getGETcaddyUrlClient().toString(),
    params,
    ...requestConfig,
    headers: { ...headers, ...requestConfig.headers },
  })
  
  return res
}
