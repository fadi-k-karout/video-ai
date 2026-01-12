<script lang="ts">
	import { onMount } from 'svelte';
	import { PUBLIC_API_BASE } from '$env/static/public';

	let apiStatus = 'Testing...';

	async function testAPI() {
		apiStatus = 'Connecting...';

		try {
			if (!PUBLIC_API_BASE) {
				throw new Error('PUBLIC_API_BASE environment variable is not set');
			}
			const apiUrl = PUBLIC_API_BASE + '/api/health';
			const response = await fetch(apiUrl);

			if (!response.ok) {
				throw new Error(`HTTP ${response.status}: ${response.statusText}`);
			}

			const data = await response.json();
			apiStatus = `✅ API Connected: ${data.message}`;
		} catch (error) {
			const errorMessage = error instanceof Error ? error.message : String(error);
			apiStatus = `❌ API Error: ${errorMessage}`;
		}
	}

	onMount(async () => {
		await testAPI();
	});
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
<p><strong>API Status:</strong> {apiStatus}</p>
<p>Check the browser console for detailed logs.</p>
