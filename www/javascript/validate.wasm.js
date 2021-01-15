function edtf_string_to_edtf_date(date){
    
    var edtf_block = document.getElementById("edtf_date_block");
    var edtf_date = document.getElementById("edtf_date");

    var feedback = document.getElementById("feedback");
    
    edtf_block.style.display = "none";
    edtf_date.innerText = "";
    
    var rsp = parse_edtf(date);

    if (! rsp){
	var item = document.createElement("li");
	item.innerText = "Failed to convert date to EDTF date with undefined (WASM) error.";
	feedback.appendChild(item);
	feedback.style.display = "block";
	return;
    }

    try {
	
	var data = JSON.parse(rsp);
	edtf_date.innerText = JSON.stringify(data, "", 2);
	edtf_block.style.display = "block";	
	
    } catch (err){	
	var item = document.createElement("li");
	item.innerText = "Unable to convert date to EDTF date: " + err;
	feedback.appendChild(item);
	feedback.style.display = "block";
	return;
    }
    
}
