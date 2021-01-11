function to_edtf_string(date){

    var edtf_block = document.getElementById("edtf_string_block");
    var edtf_str = document.getElementById("edtf_string");

    edtf_block.style.display = "none";
    edtf_str.innerText = "";
    
    var on_success = function(rsp){	
	edtf_str.innerText = rsp;
	edtf_block.style.display = "block";	
    };

    var on_error = function(err){
	console.log("ERROR", err);
    };

    sfomuseum.date.toEDTFString(date, on_success, on_error);
}

function to_edtf_date(date){

    var edtf_block = document.getElementById("edtf_date_block");
    var edtf_date = document.getElementById("edtf_date");

    edtf_block.style.display = "none";
    edtf_date.innerText = "";
    
    var on_success = function(rsp){	
	edtf_date.innerText = JSON.stringify(rsp, "", 2);
	edtf_block.style.display = "block";	
    };

    var on_error = function(err){
	console.log("ERROR", err);
    };

    sfomuseum.date.toEDTFDate(date, on_success, on_error);
}
