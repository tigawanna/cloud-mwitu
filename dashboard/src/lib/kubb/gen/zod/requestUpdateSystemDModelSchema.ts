import type { RequestUpdateSystemDModel } from '../types/RequestUpdateSystemDModel.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description RequestUpdateSystemDModel schema
 */
export const requestUpdateSystemDModelSchema = z
  .object({
    content: z.unknown().optional(),
    libDir: z.boolean().optional(),
    name: z.string().optional(),
  })
  .describe('RequestUpdateSystemDModel schema') as unknown as ToZod<RequestUpdateSystemDModel>