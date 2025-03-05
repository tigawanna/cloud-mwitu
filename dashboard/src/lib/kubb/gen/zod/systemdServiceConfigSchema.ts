import type { SystemdServiceConfig } from '../types/SystemdServiceConfig.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description SystemdServiceConfig schema
 */
export const systemdServiceConfigSchema = z.unknown() as unknown as ToZod<SystemdServiceConfig>