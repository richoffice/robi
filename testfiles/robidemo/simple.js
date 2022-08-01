; (function (input) {
    src = input[0];
    out = input[1];
    rfs = robi.Import("sample_def.json", src);
    visitors = rfs["visitors"];
    log(visitors);

    var addFunc = function(row){
        return row["name"]+"-xxx";
    };
    visitors.Add("fullname",addFunc);

    log("export to:"+out);

    robi.Export(rfs, out, "sample_def.json");
    log(visitors);
    return src;
})(input)