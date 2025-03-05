import client from '@/lib/kubb/custom-fetch-client.ts'
import type {
  GETCaddyNameQueryResponse,
  GETCaddyNamePathParams,
  GETCaddyNameHeaderParams,
  GETCaddyName400,
  GETCaddyName500,
} from "../../types/'caddyController/GETcaddyname.ts"
import type { RequestConfig, ResponseErrorConfig } from '@/lib/kubb/custom-fetch-client.ts'

export function getGETcaddynameUrlClient({ name }: { name: GETCaddyNamePathParams['name'] }) {
  return `/caddy/${name}` as const
}

/**
 * @description #### Controller: `github.com/tigawanna/cloud-mwitu/internal/controllers.CaddyFileResources.getCaddyFileServiceByName`#### Middlewares:- `github.com/go-fuego/fuego.defaultLogger.middleware`---Get caddyfile service by name
 * @summary get caddy file service by name
 * {@link /caddy/:name}
 */
export async function GETcaddynameClient(
  { name }: { name: GETCaddyNamePathParams['name'] },
  headers?: GETCaddyNameHeaderParams,
  config: Partial<RequestConfig> & { client?: typeof client } = {},
) {
  const { client: request = client, ...requestConfig } = config

  const res = await request<GETCaddyNameQueryResponse, ResponseErrorConfig<GETCaddyName400 | GETCaddyName500>, unknown>({
    method: 'GET',
    url: getGETcaddynameUrlClient({ name }).toString(),
    ...requestConfig,
    headers: { ...headers, ...requestConfig.headers },
  })
  return res
}