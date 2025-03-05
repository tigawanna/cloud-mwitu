import client from '@/lib/kubb/custom-fetch-client.ts'
import type {
  POSTSystemdMutationRequest,
  POSTSystemdMutationResponse,
  POSTSystemdHeaderParams,
  POSTSystemd400,
  POSTSystemd500,
} from "../../types/'systemdController/POSTsystemd.ts"
import type { RequestConfig, ResponseErrorConfig } from '@/lib/kubb/custom-fetch-client.ts'

export function getPOSTsystemdUrlClient() {
  return `/systemd/` as const
}

/**
 * @description #### Controller: `github.com/tigawanna/cloud-mwitu/internal/controllers.SystemDFileResources.updateSystemD`#### Middlewares:- `github.com/go-fuego/fuego.defaultLogger.middleware`---SystemDFile will be updated with matching domain record or a new one will be created
 * @summary update system d
 * {@link /systemd/}
 */
export async function POSTsystemdClient(
  data?: POSTSystemdMutationRequest,
  headers?: POSTSystemdHeaderParams,
  config: Partial<RequestConfig<POSTSystemdMutationRequest>> & { client?: typeof client } = {},
) {
  const { client: request = client, ...requestConfig } = config

  const res = await request<POSTSystemdMutationResponse, ResponseErrorConfig<POSTSystemd400 | POSTSystemd500>, POSTSystemdMutationRequest>({
    method: 'POST',
    url: getPOSTsystemdUrlClient().toString(),
    data,
    ...requestConfig,
    headers: { 'Content-Type': '*/*', ...headers, ...requestConfig.headers },
  })
  return res
}