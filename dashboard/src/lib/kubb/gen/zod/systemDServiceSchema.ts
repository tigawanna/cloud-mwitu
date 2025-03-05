import type { SystemDService } from '../types/SystemDService.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description SystemDService schema
 */
export const systemDServiceSchema = z
  .object({
    contents: z.string().optional(),
    modifiedAt: z.string().optional(),
    name: z.string().optional(),
    path: z.string().optional(),
  })
  .describe('SystemDService schema') as unknown as ToZod<SystemDService>