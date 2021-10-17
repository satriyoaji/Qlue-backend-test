let objects = JSON.parse('{"h":["username","hair_color","height"],"d":[["ali","brown",1.2],["marc","blue",1.4],["joe","brown",1.7],["zehua","black",1.8]]}');

function rearrangeJson(jsons) {
    let new_json = [];
    Object.entries(objects.d).forEach(([key, value]) => {
        new_json.push({"username" : jsons.d[key][0], "hair_color" : jsons.d[key][1], "height" : jsons.d[key][2]});
    })
    return new_json
}

console.log(rearrangeJson(objects));