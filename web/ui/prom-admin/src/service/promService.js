export default class PromService {
  getInstances() {
    return fetch('demo/data/prom-instances.json')
        .then((res) => res.json())
    }
}

// export default class PromService {
//   getInstances() {
//     return fetch('/services')
//         .then((res) => res.json())
//     }
// }
