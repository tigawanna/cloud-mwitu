import client from '@/lib/kubb/custom-fetch-client.ts'
import type {
  GETSystemdNameQueryResponse,
  GETSystemdNamePathParams,
  GETSystemdNameQueryParams,
  GETSystemdNameHeaderParams,
  GETSystemdName400,
  GETSystemdName500,
} from "../../types/'systemdController/GETsystemdname.ts"
import type { RequestConfig, ResponseErrorConfig } from '@/lib/kubb/custom-fetch-client.ts'

export function getGETsystemdnameUrlClient({ name }: { name: GETSystemdNamePathParams['name'] }) {
  return `/systemd/${name}` as const
}

/**
 * @description #### Controller: `github.com/tigawanna/cloud-mwitu/internal/controllers.SystemDFileResources.getSystemDFileServiceByName`#### Middlewares:- `github.com/go-fuego/fuego.defaultLogger.middleware`---Get SystemDFile service by name
 * @summary get system d file service by name
 * {@link /systemd/:name}
 */
export async function GETsystemdnameClient(
  { name }: { name: GETSystemdNamePathParams['name'] },
  params?: GETSystemdNameQueryParams,
  headers?: GETSystemdNameHeaderParams,
  config: Partial<RequestConfig> & { client?: typeof client } = {},
) {
  const { client: request = client, ...requestConfig } = config

  const res = await request<GETSystemdNameQueryResponse, ResponseErrorConfig<GETSystemdName400 | GETSystemdName500>, unknown>({
    method: 'GET',
    url: getGETsystemdnameUrlClient({ name }).toString(),
    params,
    ...requestConfig,
    headers: { ...headers, ...requestConfig.headers },
  })
  return res
}