import type { SuccessListResponseServicesCaddyFileModel } from '../../types/successListResponseServices/CaddyFileModel.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description SuccessListResponse_services.CaddyFileModel schema
 */
export const successListResponseServicesCaddyFileModelSchema = z
  .object({
    error: z.unknown().optional(),
    result: z
      .object({
        items: z
          .array(
            z.object({
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
            }),
          )
          .optional(),
      })
      .optional(),
  })
  .describe('SuccessListResponse_services.CaddyFileModel schema') as unknown as ToZod<SuccessListResponseServicesCaddyFileModel>