import { getEmitter } from './mitt.js'
import { fetch_get } from './fetch.js'

let emitter = getEmitter();

export default {
    name: 'Test1',

    setup() {
        const title = "Hello";
        let mypen = Vue.ref("");
        let imgsrc = Vue.ref("");

        // listen to an event
        emitter.on('from_app', e => {
            mypen.value = e;
            console.log('app1 received:', e)
        });

        function fire(str) {

            let cData = fetch_get('https://yesno.wtf/api');
            (async () => {
                const data = await cData;
                emitter.emit('from_app1', data.answer);
                console.log(data.answer);
                imgsrc.value = data.image;
            })();

            // emitter.emit('from_app1', mypen.value)
            // console.log(mypen.value);
        }

        return {
            title,
            mypen,
            fire,
            imgsrc,
        };
    },

    template: `      
        <h1>{{title}} | {{mypen}} | {{imgsrc}}</h1>
        <input v-model="mypen" placeholder="input">
        <button class="mybutton" @click="fire"></button>     
        <img :src="imgsrc" alt="YES/NO IMAGE" width="320" height="240" />  
    `,
};