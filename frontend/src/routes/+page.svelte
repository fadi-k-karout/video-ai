<script lang="ts">
	import { onMount } from 'svelte';

	console.log('Script block executed - component loading');

	let apiStatus = 'Testing...';

	// Test function we can call manually
	async function testAPI() {
		console.log('🚀 Manual API test triggered');
		apiStatus = 'Connecting...';

		try {
			console.log('Making fetch request to:', 'http://localhost:3000/api/health');
			const response = await fetch('http://localhost:3000/api/health');
			console.log('Response received:', response.status, response.statusText);

			if (!response.ok) {
				throw new Error(`HTTP ${response.status}: ${response.statusText}`);
			}

			const data = await response.json();
			console.log('✅ API Response:', data);
			apiStatus = `✅ API Connected: ${data.message}`;
		} catch (error) {
			console.error('❌ API Error:', error);
			apiStatus = `❌ API Error: ${error.message}`;
		}
	}

	console.log('About to define onMount');

	onMount(async () => {
		console.log('🚀 onMount triggered - testing API connection');
		await testAPI();
	});

	console.log('onMount defined, component script complete');
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
<p><strong>API Status:</strong> {apiStatus}</p>
<button onclick={testAPI}>Test API Manually</button>
<p>Check the browser console for detailed logs.</p>
