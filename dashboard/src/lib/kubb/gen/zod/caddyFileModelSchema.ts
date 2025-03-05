import type { CaddyFileModel } from '../types/CaddyFileModel.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description CaddyFileModel schema
 */
export const caddyFileModelSchema = z
  .object({
    block: z
      .array(
        z.object({
          path: z.array(z.string()).optional(),
          value: z.array(z.string()).optional(),
        }),
      )
      .optional(),
    content: z.string().optional(),
    domain: z.string().optional(),
    startEnd: z.unknown().optional(),
  })
  .describe('CaddyFileModel schema') as unknown as ToZod<CaddyFileModel>