<script>
	import { Grid, Column, Row, TextInput, PasswordInput } from 'carbon-components-svelte';
	import { Form, FormGroup, Button, Tile } from 'carbon-components-svelte';
	import { env } from '$env/dynamic/public';
	import { onMount, onDestroy } from 'svelte';
	const google = {
		clientId: env.PUBLIC_GOOGLE_CLIENT_ID,
		loginUri: env.PUBLIC_BASE_URL + '/api/user/auth/google',
		uxMode: 'popup',
		type: 'standard',
		locale: 'sk'
	};
	onMount(() => {
		const script = document.createElement('script');
		script.src = 'https://accounts.google.com/gsi/client';
		script.async = true;
		document.body.appendChild(script);
	});
	onDestroy(() => {
		const script = document.querySelector('script[src="https://accounts.google.com/gsi/client"]');
		if (script) {
			document.body.removeChild(script);
		}
	});
</script>

<div
	id="g_id_onload"
	data-client_id={google.clientId}
	data-login_uri={google.loginUri}
	data-ux_mode={google.uxMode}
	data-type={google.type}
	data-locale={google.locale}
	data-context="signin"
	data-itp_support={true}
/>
<Grid>
	<Row>
		<Column />
	</Row>
	<Row>
		<Column sm={0} md={2} lg={5} xlg={6} />
		<Column sm={4} md={4} lg={6} xlg={4}>
			<Tile>
				<Form>
					<FormGroup>
						<div
							class="g_id_signin"
							data-type={google.type}
							data-locale={google.locale}
							data-shape="rectangular"
							data-theme="outline"
							data-text="signin_with"
							data-size="large"
							data-logo_alignment="left"
							data-width="100%"
						/>
					</FormGroup>
					<FormGroup>
						<TextInput labelText="User name" placeholder="Enter user name..." />
					</FormGroup>
					<FormGroup>
						<PasswordInput labelText="Password" placeholder="Enter password..." />
					</FormGroup>
					<FormGroup>
						<Button type="submit">Sign In</Button>
					</FormGroup>
				</Form>
			</Tile>
		</Column>
		<Column sm={0} md={2} lg={5} xlg={6} />
	</Row>
</Grid>
