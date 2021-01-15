function edtf_string_to_edtf_date(date){

    var edtf_block = document.getElementById("edtf_date_block");
    var edtf_date = document.getElementById("edtf_date");

    var feedback = document.getElementById("feedback");
    
    edtf_block.style.display = "none";
    edtf_date.innerText = "";

    var on_success = function(rsp){
	edtf_date.innerText = JSON.stringify(rsp, "", 2);
	edtf_block.style.display = "block";	
    };

    var on_error = function(err){
	var item = document.createElement("li");
	item.innerText = "Unable to convert date to EDTF date: " + err;
	feedback.appendChild(item);
	feedback.style.display = "block";
    };

    sfomuseum.date.edtfStringtoEDTFDate(date, on_success, on_error);
}
