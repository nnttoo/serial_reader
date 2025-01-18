import { createApp } from 'vue'
import App from './App.vue'
import './style.css';

function setupLinkRedirect(){
    document.body.addEventListener('click', function (e: any) {
        if (e.target && e.target.nodeName == 'A' && e.target.href) {
            const url = e.target.href;
            if (
                !url.startsWith('http://#') &&
                !url.startsWith('file://') &&
                !url.startsWith('http://wails.localhost:')
            ) {
                e.preventDefault();
                (window as any).runtime.BrowserOpenURL(url);
            }
        }
    });
}

setupLinkRedirect();
createApp(App).mount('#app')
