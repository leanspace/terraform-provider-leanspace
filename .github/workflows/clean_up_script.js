module.exports = async (service_list, tenant, client_id, client_secret) => {
	async function delete_resource(endpoint, headers) {
		console.log(endpoint);
		const delete_service = await fetch(endpoint, {
			method: "DELETE",
			headers
		});
		await delete_service.text();
		if (delete_service.status !== 200) {
			console.log(`Failed to delete ${delete_service.json()}`);
		}
	}
	const token_request = await fetch(`https://api.develop.leanspace.io/teams-repository/oauth2/token?tenant=${tenant}`, {
		method: "POST",
		headers: {
			"Content-type": "application/x-www-form-urlencoded",
			"Authorization": `Basic ${Buffer.from(`${client_id}:${client_secret}`).toString('base64')}`
		}
	});

	const token = (await token_request.json()).access_token;
	const headers = {
		"Content-type": "application/json; charset=UTF-8",
		"Authorization": `Bearer ${token}`
	};
	const services = service_list.split(',');

	for (const service of services) {
		const endpoint = encodeURI(`https://api.develop.leanspace.io/${service}`);
		const get_all = await fetch(endpoint, {
		method: "GET",
		headers
		});
		const all = service.includes('?') ? (await get_all.json()).content : await get_all.json();
		if (get_all.status !== 200 ) {
			console.log(`Failed to get ${service}`);
			continue;
		}
		console.log(all);
		console.log(endpoint);
		if (service.includes('?')) {
			for (const item of all) {
				if (item.name === 'terraform auto checker') {
					console.log("Ignoring terraform auto checker");
					continue;
				}
				const get_all_endpoint = encodeURI(`https://api.develop.leanspace.io/${service.split('?')[0]}/${item.id}`);
				delete_resource(get_all_endpoint, headers);
			}
		} else {
			delete_resource(endpoint, headers);
		}
	}
}