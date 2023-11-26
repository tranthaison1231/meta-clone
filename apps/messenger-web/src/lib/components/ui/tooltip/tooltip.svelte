<script lang="ts">
	import { Tooltip as TooltipPrimitive } from 'bits-ui';
	import { cn } from '$lib/utils/cn';
	import { flyAndScale } from '$lib/utils/flyAndScale';

	export let content = '';
	export let className: string | undefined = undefined;
</script>

<TooltipPrimitive.Root>
	<TooltipPrimitive.Trigger>
		<slot />
	</TooltipPrimitive.Trigger>
	<TooltipPrimitive.Content
		transition={flyAndScale}
		transitionConfig={{
			y: 8,
			duration: 150
		}}
		sideOffset={4}
		class={cn(
			'z-50 overflow-hidden rounded-md border bg-popover px-3 py-1.5 text-sm text-popover-foreground shadow-md',
			className
		)}
		{...$$restProps}
	>
		{content}
	</TooltipPrimitive.Content>
</TooltipPrimitive.Root>
