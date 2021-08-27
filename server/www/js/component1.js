import { getEmitter } from './mitt.js'
import { fetch_get_json } from './fetch.js'

let emitter = getEmitter();

export default {
    name: 'Test',

    setup() {
        const title = "Hello";
        let mypen = Vue.ref("");
        let imgsrc = Vue.ref("");
        let timer_str = Vue.ref("");

        // listen to an event
        emitter.on('from_app1', e => {
            mypen.value = e;
            console.log('app received:', e);
        });

        // listen to all events
        // emitter.on('*', (type, e) => console.log(type, e));

        // fire an event
        // emitter.emit('foo', { a: 'b' })

        function fire() {

            let cData = fetch_get_json('https://yesno.wtf/api');

            // 'async function' return channel             
            const fnFetchValue = async () => {
                const data = await cData;
                emitter.emit('from_app', data.answer);
                console.log(data.answer);
                imgsrc.value = data.image;
                return data;
            };
            // 'async function' return channel
            let cOut = fnFetchValue();
            console.log(`com1 result is ${cOut}`)

            // emitter.emit('from_app', mypen.value);
            // console.log(mypen.value);
        }

        let myTimer = setInterval(
            () => { timer_str.value = (new Date()).toLocaleTimeString(); },
            1000,
        )
        // clearInterval(myTimer);

        return {
            title,
            mypen,
            fire,
            imgsrc,
            timer_str,
        };
    },

    template: `      
        <h1>{{title}} | {{timer_str}} | {{mypen}} | {{imgsrc}}</h1>
        <input v-model="mypen" placeholder="input">
        <button class="mybutton" @click="fire"></button>   
        <img :src="imgsrc" alt="YES/NO IMAGE" width="320" height="240" />   
    `,
};