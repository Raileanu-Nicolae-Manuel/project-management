import { createSignal, For } from "solid-js";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "../ui/card";
import AsyncButton from "./async-button";
import { await1000 } from "@/api/fake";


const CardDemo = () => {
	return (
		<Card class="w-[380px]">
			<CardHeader>
				<CardTitle>test</CardTitle>
				<CardDescription>test</CardDescription>
			</CardHeader>
			<CardContent class="grid gap-4">
				test
			</CardContent>
			<CardFooter>
				<AsyncButton class="w-full" onClick={await1000}>
					Time out button
				</AsyncButton>
			</CardFooter>
		</Card>
	);
};

export default CardDemo;