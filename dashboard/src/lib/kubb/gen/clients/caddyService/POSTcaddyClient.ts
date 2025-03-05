import client from '@/lib/kubb/custom-fetch-client.ts'
import type {
  POSTCaddyMutationRequest,
  POSTCaddyMutationResponse,
  POSTCaddyHeaderParams,
  POSTCaddy400,
  POSTCaddy500,
} from "../../types/'caddyController/POSTcaddy.ts"
import type { RequestConfig, ResponseErrorConfig } from '@/lib/kubb/custom-fetch-client.ts'

export function getPOSTcaddyUrlClient() {
  return `/caddy/` as const
}

/**
 * @description #### Controller: `github.com/tigawanna/cloud-mwitu/internal/controllers.CaddyFileResources.updateCaddy`#### Middlewares:- `github.com/go-fuego/fuego.defaultLogger.middleware`---Caddyfile will be updated with matching domain record or a new one will be created
 * @summary update caddy
 * {@link /caddy/}
 */
export async function POSTcaddyClient(
  data?: POSTCaddyMutationRequest,
  headers?: POSTCaddyHeaderParams,
  config: Partial<RequestConfig<POSTCaddyMutationRequest>> & { client?: typeof client } = {},
) {
  const { client: request = client, ...requestConfig } = config

  const res = await request<POSTCaddyMutationResponse, ResponseErrorConfig<POSTCaddy400 | POSTCaddy500>, POSTCaddyMutationRequest>({
    method: 'POST',
    url: getPOSTcaddyUrlClient().toString(),
    data,
    ...requestConfig,
    headers: { 'Content-Type': '*/*', ...headers, ...requestConfig.headers },
  })
  return res
}