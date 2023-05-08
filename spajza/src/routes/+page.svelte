<script>
	// redirect to login page if not logged in
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	onMount(() => {
		goto('/login');
	});
</script>
