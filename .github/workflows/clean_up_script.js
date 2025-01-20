module.exports = async (tenant, env, client_id, client_secret) => {
	const token_request = await fetch(`https://api.develop.leanspace.io/teams-repository/oauth2/token?tenant=${tenant}`, {
		method: "POST",
		body: `client_id=${client_id}&client_secret=${client_secret}&grant_type=client_credentials`,
		headers: {
		"Content-type": "application/x-www-form-urlencoded",
		}
	});

	const token = (await token_request.json()).access_token;
	const headers = {
		"Content-type": "application/json; charset=UTF-8",
		"Authorization": `Bearer ${token}`
	};
	const services = ['activities-repository/activity-definitions?query=Terraform%','activities-repository/activities/states?query=MY_TEST','activities-repository/activity-definitions/resource-functions?query=Terraform%',
		'agents-repository/agents?query=Terraform%',
		`asset-repository/nodes?createdBy=${client_id}`, `asset-repository/properties/v2?createdBy=${client_id}`, 'asset-repository/units?query=%Customium',
		"commands-repository/command-definitions?createdBys=${createdBy}", 'commands-repository/command-queues?query=Terraform%', 'commands-repository/command-sequences/states?query=MY_TEST', 'commands-repository/command-sequences/commands/states?query=MY_TEST',`commands-repository/release-queues?createdBys=${client_id}`,
		'dashboard-repository/dashboards?query=Terraform%', 'dashboard-repository/widgets?query=Terraform%',
		'events/event-definitions?query=Terraform%',
		'integration-leafspace/connections', 'integration-leafspace/contact-reservations/status/mappings?query=terraform test', 'integration-leafspace/ground-stations/links??leafspaceGroundStationIds=d5de2269dc23179929546f41b6239afb', 'integration-leafspace/satellites/links?leafspaceSatelliteIds=c736e7ed36916f5d4c14e454087a3dc2',
		'metrics-repository/metrics?query=Terra%',
		'monitors-repository/action-templates?query=Terraform&','monitors-repository/monitors?query=Terraform%',
		`orbits-repository/orbits?createdBys=${client_id}`,
		'passes-repository/contacts/states?query=MY_TEST', 'passes-repository/passes/delay/configurations', 'passes-repository/passes/states?query=MY_TEST',
		'plans-repository/plans/states?query=MY_TEST', `plans-repository/plan-templates?createdBys=${client_id}`,
		'plugins-repository/generic-plugins?query=Terraform%', 'plugins-repository/plugins?query=Terraform%',
		'records/record-templates?query=Terraform%',
		`requests-repository/feasibility-constraint-definitions?createdBys=${client_id}`, 'requests-repository/request-definitions?query=%Terraform', 'requests-repository/requests/states?query=TERRAFORM_STATE',
		`resources-repository/resources?createdBys=${client_id}`,
		`routes-repository/processors?createdBys=${client_id}`, 'routes-repository/routes?query=Terraform%',
		'streams-repository/streams?query=Terraform%',
		'teams-repository/access-policies?query=Terraform%', 'teams-repository/members?query=Terra%', 'teams-repository/service-accounts?query=Terra%', 'teams-repository/teams?query=Terra%',
	];

	for (const service of services) {
		const get_all = await fetch(encodeURI(`https://api.develop.leanspace.io/${service}`), {
		method: "GET",
		headers
		});
		const all = service.includes('?') ? (await get_all.json()).content : await get_all.json();
		if (get_all.status !== 200 ) {
			console.log(`Failed to get ${service}`);
			continue;
		}
		console.log(all);
		console.log(`https://api.develop.leanspace.io/${service}`);
		if (service.includes('?')) {
			for (const item of all) {
				if (item.name === 'terraform auto checker') {
					console.log("Ignoring terraform auto checker");
					continue;
				}
				console.log(`https://api.develop.leanspace.io/${service.split('?')[0]}/${item.id}`);
				const delete_service = await fetch(encodeURI(`https://api.develop.leanspace.io/${service.split('?')[0]}/${item.id}`), {
					method: "DELETE",
					headers
				});
				await delete_service.text();
				if (delete_service.status !== 200) {
					console.log(`Failed to delete ${delete_service.json()}`);
				}
			}
		} else {
			console.log(`https://api.develop.leanspace.io/${service}`);
			const delete_service = await fetch(`https://api.develop.leanspace.io/${service}`, {
				method: "DELETE",
				headers
			});
			await delete_service.text();
			if (delete_service.status !== 200) {
				console.log(`Failed to delete ${delete_service.json()}`);
			}
		}
	}
}