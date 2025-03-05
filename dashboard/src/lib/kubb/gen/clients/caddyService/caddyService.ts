import { GETcaddyClient } from './GETcaddyClient.ts'
import { GETcaddynameClient } from './GETcaddynameClient.ts'
import { POSTcaddyClient } from './POSTcaddyClient.ts'

export function caddyService() {
  return { GETcaddyClient, POSTcaddyClient, GETcaddynameClient }
}