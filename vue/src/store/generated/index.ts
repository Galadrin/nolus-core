// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import NomoCosmzoneNolusMintV1Beta1 from './nomo/cosmzone/nolus.mint.v1beta1'
import NomoCosmzoneNomoCosmzoneTax from './nomo/cosmzone/nomo.cosmzone.tax'


export default { 
  NomoCosmzoneNolusMintV1Beta1: load(NomoCosmzoneNolusMintV1Beta1, 'nolus.mint.v1beta1'),
  NomoCosmzoneNomoCosmzoneTax: load(NomoCosmzoneNomoCosmzoneTax, 'nomo.cosmzone.tax'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
