import { integer, timestamp } from "drizzle-orm/pg-core";
export const defaultTable = {
    id: integer().primaryKey().generatedAlwaysAsIdentity(),
    createdAt: timestamp().defaultNow(),
    updatedAt: timestamp().defaultNow(),
};