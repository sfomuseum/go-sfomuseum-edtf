window.addEventListener("load", function load(event){

    var submit_button = document.getElementById("submit");

    if (! submit_button){
	console.log("Missing submit button");
	return;
    }
    
    if (! WebAssembly.instantiateStreaming){
	
	WebAssembly.instantiateStreaming = async (resp, importObject) => {
            const source = await (await resp).arrayBuffer();
            return await WebAssembly.instantiate(source, importObject);
	};
    }
    
    const go = new Go();
    
    let mod, inst;
    
    WebAssembly.instantiateStreaming(fetch("../static/wasm/parse.wasm"), go.importObject).then(
	
	async result => {

	    submit_button.onclick = function(){
		
		var date_el = document.getElementById("date");
		
		if (! date_el){
		    console.log("Missing date el");
		    return false;
		}
		
		var date_str = date_el.value;
		
		if (date_str == ""){
		    console.log("Empty date string");
		    return false;
		}
		
		var feedback = document.getElementById("feedback");
		feedback.style.display = "none";	
		feedback.innerHTML = "";

		try {
		    edtf_string_to_edtf_date(date_str);
		} catch (err){
		    var item = document.createElement("li");
		    item.innerText = "Unable to convert date to EDTF date: " + err;
		    feedback.appendChild(item);
		    feedback.style.display = "block";
		}
		
		return false;
	    };

	    submit_button.removeAttribute("disabled");	    
	    
            mod = result.module;
            inst = result.instance;
	    await go.run(inst);
	}
    );
    
});
