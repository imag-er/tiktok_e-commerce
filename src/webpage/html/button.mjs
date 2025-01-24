// 按钮样式变化逻辑，不包含请求逻辑

import {apis, fetchApiData} from './request.mjs';

const apiList = document.getElementById('api-list');

function updateButtonStyle(button, response) {
    const responseDiv = button.nextElementSibling;
    responseDiv.style.display = 'block';
    responseDiv.innerHTML = '<pre>' + JSON.stringify(response, null, 2) + '</pre>';
}

function handleError(button, error) {
    const responseDiv = button.nextElementSibling;
    responseDiv.style.display = 'block';
    responseDiv.innerHTML = 'Error: ' + error.message;
}

function testApi(name, button) {
    updateButtonStyle(button, 'Loading...');
    fetchApiData(name)
        .then(data => {
            updateButtonStyle(button, data);
        })
        .catch(error => {
            handleError(button, error);
        });
}

Object.keys(apis).forEach(apiName => {
    const apiDiv = document.createElement('div');
    apiDiv.className = 'api';

    const apiTitle = document.createElement('h3');
    apiTitle.textContent = apiName;
    apiDiv.appendChild(apiTitle);

    const apiButton = document.createElement('button');
    apiButton.textContent = 'Test Connection';
    apiButton.onclick = () => testApi(apiName, apiButton);
    apiDiv.appendChild(apiButton);

    const apiResponseDiv = document.createElement('div');
    apiResponseDiv.className = 'response';
    apiResponseDiv.style.display = 'none';
    apiDiv.appendChild(apiResponseDiv);

    apiList.appendChild(apiDiv);
});




export {testApi}