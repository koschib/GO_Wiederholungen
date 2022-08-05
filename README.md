# GO_Wiederholungen

Aufgabe I (Golang Basics) Theorie zu den Basics
1. Welche Datentypen gibt es Golang nicht!
Int; int8; int16; int32; int64; int128; float; float16; float32; float64 rune; char; byte; String; string; boolean;
2. Bezeichnen Sie die einzelnen Teile der vorliegenden Go-Datei und erklären Sie auch deren Eigenschaften!
3. Übersetzen Sie das nachfolgende Java Programm in Golang Code!
4. Warum ist Golang gut geeignet für die Entwicklung von Microservices? 5. Wo werden Microservices eingesetzt?
   
Aufgabe II (Go in a Nutshell)
1. Welche Arten zur Deklaration und Initialisierung von Variablen kennen Sie in Go?
2. Vervollständigen Sie die nachfolgende Go-Datei. Füllen Sie nur die gekennzeichneten Stellen aus und fügen keinen weiteren Code hinzu.
S
3. Erklären Sie den Unterschied zwischen Arrays in Golang und Slices!
4. Erstellen Sie ein Array des Typen „int“ mit dem Namen „Array1“. Das Array soll eine Länge von 10 haben. Startwert an der Index-Stelle 0 ist 2. Jede weitere Indexstelle soll um den doppelten Wert erhöht werden.
• Filtern Sie die Werte von Index 4 bis 9 mit Slice und bezeichnen diese als „slice1“!
• Was ist die Kapazität und was die Länge von „slice1“?
• Rechnen Sie die Index Inhalte aus Ihrem „slice1“ zusammen und ermitteln so die
Summe.
• Geben Sie ein Befehl an, mit dem Sie ein Element des Int-Arrays „Array1“ abfragen
können.
• Geben Sie ein Befehl an, mit dem Sie ein Element des int-Arrays „Array1“ ändern
können.
• Erstellen Sie aus demselben int-Array „Array1“ ein Slice mit der Funktion make(). Der
Bereich sollte mindestens über 5 Index Werte gehen. Das neue Slice soll „slice2“
heißen.
• Geben Sie ein Befehl an, mit dem Sie ein Element von „slice2“ abfragen
• Geben Sie ein Befehl an, mit dem Sie Elemente von „slice2“
• Erweitern Sie ihr „slice2“ um mindestens drei weitere Elemente.
  
5. Schreiben Sie die Syntax einer Funktion in der Grundstruktur mit einem Rückgabewert und einer Funktion ohne Rückgabewert. Was sind die Unterschiede und wozu gibt es überhaupt Rückgabewerte?
6. Schreiben Sie eine rekursive Funktion in Go. Die Funktion soll einen Wert „x“ solange mit einem Wert y vergleichen bis dieser übereinstimmt. Nach jedem rekursiven Aufruf soll die Ausgabe „Stimmt noch nicht, noch eine Runde!“ ausgegeben werden. Stimmen die Werte überein soll die Ausgabe „Jetzt passt das. Wir sind fertig!“ ausgegeben werden.
Aufgabe II (Go Advanced)
1. Theorie Fragen!
Wie wird das Verzeichnis im Strukturbaum genannt in dem go-Dateien meistens abgelegt werden?
In welchem Ordner landen Binärdateien standardmäßig, wenn sie mit dem Befehl “go install” installiert werden?
Was kommt in der Go-Syntax zuerst? Der Variablenname oder die Deklaration des Typs?
2. Erstelle eine mehrfache if/else-Ausgabe mit folgendem Aufbau und Bedingungen:
-Wenn Eingabe:
-größer 60, print: “Die Person ist alt geworden.”
-größer 30, print: “Die Person ist weiser geworden.” -größer 20, print: “Die Person ist erwachsen geworden.” -größer 10, print: “Die Person ist noch jung.”
-sonst, print: “Wächst noch auf.”
 
3. Dein Chef will, das du ein Programm entwickelst, das prüft ob eine Person einen Film schauen darf oder nicht.
Das Programm soll prüfen welches Alter eingegeben wurde und je nachdem andere Texte ausgeben.
1. Rufen Sie das Alter von der Befehlszeile ab.
2. Geben Sie eine der folgenden Nachrichten zurück, wenn das Alter lautet:
-> Über 17: "R-Rated"
-> Zwischen 13 und 17: "PG-13" -> Unter 13 : "PG-bewertet"
BESCHRÄNKUNGEN:
1. Wenn die Altersangaben falsch sind oder fehlen, teilen Sie dies dem Benutzer mit. 2. Akzeptieren Sie kein negatives Alter.
BONUS:
1. BONUS: Verwenden Sie if-Anweisungen in Ihrem Programm nur zweimal. 2. BONUS: Verwenden Sie eine if-Anweisung nur einmal.
4. Erstelle vier switch cases mit folgenden Unterscheidungen:
1. Eingabe “i”
3 Cases:
-Wenn i = 1, print “Eins.” -Wenn i = 2, print “Zwei.” -Wenn i = 3, print “Drei.”
2. Unterscheidung Wochentag oder Wochenende:
Eingabe “Wochentag”
2 Cases:
-Wenn Weekday = Montag-Freitag: “Es ist eine Wochentag” -Wenn Weekday = Samstag & Sonntag: “Es ist Wochenende”
3. Unterscheidung Vormittag/Nachmittag: Eingabe “Jetzt”
2 Cases:
-Wenn Now = 0-12: “Es ist Vormittags.” -Wenn Now = 13-24: “Es ist Nachmittags.”
4. Schreibe eine Funktion die unterscheidet ob die Eingabe “boolean”, “int” oder nichts von beiden ist und dies jeweils in einem kurzen Satz ausgibt.
 
5. Vervollständigen Sie die nachfolgende Go-Datei. Nehmen Sie keine Änderungen an den gegebenen Werten vor. Nutzen Sie alle und ausschließlich nur die Packages die Sie im Import-Teil vorgegeben bekommen haben.
  
