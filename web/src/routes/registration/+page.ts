// since there's no dynamic data here, we can prerender
// it so that it gets served as a static asset in production
export const prerender = true;

const authenticated = true;
const authorized = true;

// Only allow if authenticated
function doX() {}

// function main() {
//     if(authenticated){
//         doY();
//     }else{
//         doX();
//     }
// }

function main() {
    if (!authenticated || !authorized) return;

    doX();
}
