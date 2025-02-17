// 请求逻辑，不包含任何DOM操作


export const apis = {
    hello: {
        url: 'http://127.0.0.1:8888/hello',
        method: "GET",
        arguments: {},
        contentType: 'none' // No content needed for GET
    },
    login: {
        url: 'http://127.0.0.1:8888/login',
        method: "POST",
        arguments: { email: "admin", password: "admin" },
        contentType: 'form-data'
    },
    updateUser: {
        url: 'http://127.0.0.1:8888/api/user/update',
        method: "PUT",
        arguments: { name: "John Doe", age: 30 },
        contentType: 'json'
    },
    deleteUser: {
        url: 'http://127.0.0.1:8888/api/user/delete',
        method: "DELETE",
        arguments: { userId: 123 },
        contentType: 'param'
    },
};

export function fetchApiData(apiName) {
    const api = apis[apiName];
    const fetchArgus = {
        method: api.method,
    };

    switch (api.contentType) {
        case 'json':
            fetchArgus.body = JSON.stringify(api.arguments);
            fetchArgus.headers = {
                'Content-Type': 'application/json',
            };
            break;
        case 'form-data':
            let formData = new FormData();
            for (const key in api.arguments) {
                formData.append(key, api.arguments[key]);
            }
            fetchArgus.body = formData;
            fetchArgus.headers = {
                'Content-Type': 'multipart/form-data',
            };
            break;
        case 'url-encoded':
            let urlEncodedData = new URLSearchParams();
            for (const key in api.arguments) {
                urlEncodedData.append(key, api.arguments[key]);
            }
            fetchArgus.body = urlEncodedData;
            fetchArgus.headers = {
                'Content-Type': 'application/x-www-form-urlencoded',
            };
            break;
        case 'param':
            let params = new URLSearchParams(api.arguments).toString();
            api.url += '?' + params;
            break;
        case 'none':
            // No content to be sent
            break;
    }

    return fetch(api.url, fetchArgus)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok ' + response.statusText);
            }
            return response.json();
        });
}
