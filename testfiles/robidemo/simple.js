; (function (input) {
    src = input[0];
    out = input[1];
    rfs = robi.Import("sample_def.json", src);
    visitors = rfs["visitors"];
    log(visitors);

    var addFunc = function(row){
        return row["name"]+"-xxx";
    };
    visitors.Mutate("fullname",addFunc);

    log("export to:"+out);

    robi.Export(rfs,"sample_def.json", out );
    return src;
})(input)