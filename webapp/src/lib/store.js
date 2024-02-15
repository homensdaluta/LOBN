import { writable } from 'svelte/store';

export const inetName=writable('')

const inet = writable({
    name: '',
    description: '',
    flags: '',
    addresses: Object.freeze([{
        IP: '',
        Netmask: ''
    }])
});

export function setInet(name){
    inet.update(
        $inet => $inet.name = name
    );
}