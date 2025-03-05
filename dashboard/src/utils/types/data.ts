// import { GetApiAuthMe200 } from "../../lib/kubb/gen";


export type ErrorSchema = {
    message: string;
    code?: string;
    data?: Record<string, {
        code: string;
        message: string;
    }> | undefined;
}

export interface ListResultSchema<T extends Record<string,any> = Record<string,any>>{
    page: number;
    perPage: number;
    totalItems: number;
    totalPages: number;
    items: Array<T>;
}

type GenericResponse<T=any> = {
    result: T;
    error?: ErrorSchema;
}

export type InventoryUser= GenericResponse["result"]
