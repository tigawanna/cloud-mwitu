import client from '@/lib/kubb/custom-fetch-client.ts'
import type {
  GETSystemdQueryResponse,
  GETSystemdQueryParams,
  GETSystemdHeaderParams,
  GETSystemd400,
  GETSystemd500,
} from "../../types/'systemdController/GETsystemd.ts"
import type { RequestConfig, ResponseErrorConfig } from '@/lib/kubb/custom-fetch-client.ts'

export function getGETsystemdUrlClient() {
  return `/systemd/` as const
}

/**
 * @description #### Controller: `github.com/tigawanna/cloud-mwitu/internal/controllers.SystemDFileResources.getSystemDFileServices`#### Middlewares:- `github.com/go-fuego/fuego.defaultLogger.middleware`---List all SystemDFile services and filter by name
 * @summary get system d file services
 * {@link /systemd/}
 */
export async function GETsystemdClient(
  params?: GETSystemdQueryParams,
  headers?: GETSystemdHeaderParams,
  config: Partial<RequestConfig> & { client?: typeof client } = {},
) {
  const { client: request = client, ...requestConfig } = config

  const res = await request<GETSystemdQueryResponse, ResponseErrorConfig<GETSystemd400 | GETSystemd500>, unknown>({
    method: 'GET',
    url: getGETsystemdUrlClient().toString(),
    params,
    ...requestConfig,
    headers: { ...headers, ...requestConfig.headers },
  })
  return res
}