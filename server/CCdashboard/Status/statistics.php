<?php

### TAKE the Numbers
//$createinfect = file_get_contents("MIDDLEWAREXXX/dashboard/status/createinfect");
$createinfect = '{"created":280,"infected":80}';
$createinfect = json_decode($createinfect,true);
//$types_count = file_get_contents("MIDDLEWAREXXX/dashboard/status/typescount");
$types_count = '{"phishing":200,"adware":50,"wulnware":30}';
$types_count = json_decode($types_count,true);
//$payments = file_get_contents("MIDDLEWAREXXX/dashboard/status/paid");
$payments = '{"paid":50,"notpaid":30,"collected":3000,"target":10000}';
$payments = json_decode($payments,true);
//$mails = file_get_contents("MIDDLEWAREXXX/dashboard/status/mails");
$mails = '{"sent":200,"opened":125,"download":70,"infectedbyphishingmail":30}';
$mails = json_decode($mails,true);






function createDonutDiv($id,$title){
    $htmldiv ='<div class="col-md-5 py-1">
                <h2 class="display-5">'.$title.'</h2>
                <div class="card">
                    <div class="card-body">
                        <canvas id="'.$id.'"></canvas>
                    </div>
                </div>
            </div>';
    return $htmldiv;
}

?>


<!-- HTML -->

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>APT_Al</title>
    <!-- CSS -->
    <link rel="stylesheet" href="/css/style.css">
    <link rel="stylesheet" href="/css/bootstrap.min.css">

</head>
    <body>

        <div class="topnav">
            <nav class="navbar navbar-expand-lg navbar-dark" style="background-color: #070812">
                <a class="navbar-brand" href="/Status"><img src="/image/logo1.png" style="height: 60px; width: auto"></a>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>

                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav mr-auto">
                        <li class="nav-item">
                            <a class="nav-link" href="/Status/index.php">STATS</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="/Payment/index.php">PAYMENT</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="/Botnet/index.php">BOTNET</a>
                        </li>
                    </ul>
                </div>
            </nav>
        </div>

        <div class="container">
            <div class="row">
                <div class="sidenav" style="margin-top: 86px; width: auto; background-color: #242640">
                    <a href="index.php">Active Victims</a><!-- index page -->
                    <a href="statistics.php">Statics</a> <!-- it will produced using payment rate, amount of collected money... -->
                </div>
                <div class="col-sm">

                    <script src="/js/chart.js"></script>


                    <div class="container" id="graphics">
                        <h1 class="display-5">HEADER1</h1>
                            <div class="row py-2">
                                <?php
                                
                                    $titles = array("Target Amount - Collected Money","Created - Infected", "Types", "Paid - Not Paid", "Sent Mail - Opened Mail",
                                                    "Opened Mail - Infected");
                                    $donut_count = 0;
                                    foreach($titles as $tit){
                                        echo createDonutDiv("chDonut".++$donut_count,$tit);
                                    }

                                ?>

                        
                        </div>
                    </div>

                    <script>
                        var colors = ['#ff0000','#007bff','#c3e6cb','#28a745','#333333','#dc3545','#6c757d'];

                        function donutmaker(elementid, labels, data){
                            
                            var donutOptions = {
                            cutoutPercentage: 55, 
                            legend: {position:'bottom', padding:5, labels: {pointStyle:'circle', usePointStyle:true}}
                            };

                            
                            var chDonutData = {
                                labels: labels,
                                datasets: [
                                {
                                    backgroundColor: colors.slice(0,2),
                                    borderWidth: 0,
                                    data: data
                                }
                                ]
                            };

                            var chDonut = document.getElementById(elementid);
                            if (chDonut) {
                                var cd=new Chart(chDonut, {
                                    type: 'pie',
                                    data: chDonutData,
                                    options: donutOptions
                                });
                            }

                        }

                        <?php        

                            $slice_names = array();
                            $slice_values = array();
                            foreach($titles as $tit){
                                if($tit === "Types"){$tit = "phishing - adware - wulnware";}
                                array_push($slice_names, explode(" - ",$tit));
                            }
                            $slice_values = array(
                                array($payments["target"],$payments["collected"]),
                                array($createinfect["created"],$createinfect["infected"]),
                                array($types_count["phishing"],$types_count["adware"],$types_count["wulnware"]),
                                array($payments["paid"],$payments["notpaid"]),
                                array($mails["sent"],$mails["opened"]),
                                array($mails["download"],$mails["infectedbyphishingmail"]),
                            );
                        
                            $slice_names = json_encode($slice_names);
                            $slice_values = json_encode($slice_values);
                            echo "var slice_names = $slice_names;";
                            echo "var slice_values = $slice_values;";
                            echo "var donut_count= $donut_count;";
                        ?>
                        
                        for (var i = 1; i<=donut_count; i++){   
                            donutmaker("chDonut"+i,slice_names[i-1],slice_values[i-1]);
                        }
                        

                    </script>
                </div>
            </div>
        </div>

        <script src="/js/jquery-3.5.1.slim.min.js"></script>
        <script src="/js/bootstrap.bundle.min.js"></script>

    </body>
</html>