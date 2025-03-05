import client from '@/lib/kubb/custom-fetch-client.ts'
import type {
  GETSystemdRunningQueryResponse,
  GETSystemdRunningQueryParams,
  GETSystemdRunningHeaderParams,
  GETSystemdRunning400,
  GETSystemdRunning500,
} from "../../types/'systemdController/GETsystemdrunning.ts"
import type { RequestConfig, ResponseErrorConfig } from '@/lib/kubb/custom-fetch-client.ts'

export function getGETsystemdrunningUrlClient() {
  return `/systemd/running` as const
}

/**
 * @description #### Controller: `github.com/tigawanna/cloud-mwitu/internal/controllers.SystemDFileResources.getRunningSystemDFileServices`#### Middlewares:- `github.com/go-fuego/fuego.defaultLogger.middleware`---List all Running SystemDFile services and filter by name
 * @summary get running system d file services
 * {@link /systemd/running}
 */
export async function GETsystemdrunningClient(
  params?: GETSystemdRunningQueryParams,
  headers?: GETSystemdRunningHeaderParams,
  config: Partial<RequestConfig> & { client?: typeof client } = {},
) {
  const { client: request = client, ...requestConfig } = config

  const res = await request<GETSystemdRunningQueryResponse, ResponseErrorConfig<GETSystemdRunning400 | GETSystemdRunning500>, unknown>({
    method: 'GET',
    url: getGETsystemdrunningUrlClient().toString(),
    params,
    ...requestConfig,
    headers: { ...headers, ...requestConfig.headers },
  })
  return res
}