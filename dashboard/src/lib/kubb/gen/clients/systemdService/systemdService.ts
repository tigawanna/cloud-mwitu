import { GETsystemdClient } from './GETsystemdClient.ts'
import { GETsystemdnameClient } from './GETsystemdnameClient.ts'
import { GETsystemdrunningClient } from './GETsystemdrunningClient.ts'
import { POSTsystemdClient } from './POSTsystemdClient.ts'

export function systemdService() {
  return { GETsystemdClient, POSTsystemdClient, GETsystemdrunningClient, GETsystemdnameClient }
}