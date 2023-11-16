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

    // See this: We are probably not loading this from the same bundle that this file
    // (validate.wasm.init.js) is included with. That's because this is currently run
    // inside of a Lambda function and Lambda won't serve the .wasm file complaining
    // the "body (of the file) is too big". Which is weird because the WASM file is ~3MB
    // and Lambda has a reported limit of 6MB so I don't know what's going on. Instead
    // this is being served directly from the MF website. Computers, amirite...
    // (20210115/thisisaaronland)

    // WebAssembly.instantiateStreaming(fetch("../static/javascript/parse.wasm"), go.importObject).then(

    // So instead we just read the WASM file hosted from the millsfield.sfomuseum.org website. Womp womp...
    WebAssembly.instantiateStreaming(fetch("/wasm/edtf/parse.wasm"), go.importObject).then(	
	
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

		var result_el = document.getElementById("edtf_date_block");
		
		parse_edtf(date_str).then(rsp => {
		    
		    try {
			var edtf_d = JSON.parse(rsp)
		    } catch(e){
			result_el.innerText = "Unable to parse your EDTF string: " + e;
			
			result_el.style.display = "block";
			return;
		    }
		    
		    var pre = document.createElement("pre");
		    pre.innerText = JSON.stringify(edtf_d, '', 2);
		    
		    result_el.appendChild(pre);
		    result_el.style.display = "block";
		    
		}).catch(err => {
		    result_el.innerText = "There was a problem parsing your EDTF string:" + err;
		    result_el.style.display = "block";
		});
		
		return false;
	    };

	    submit_button.innerText = "Validate";
	    submit_button.removeAttribute("disabled");	    
	    
            mod = result.module;
            inst = result.instance;
	    await go.run(inst);
	}
    );
    
});
