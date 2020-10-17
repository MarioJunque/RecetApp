
JFLAGS = -g
JC = javac
JAR = jar
JRUN = java
.SUFFIXES: .java .class

.java.class:
	$(JC) $(JFLAGS) $*.java


MANIFEST = app
CLASSES = \
        Foo.java \
        Hello.java

default: classes

classes: $(CLASSES:.java=.class)

clean:
	$(RM) *.class *.jar

jar: $(MANIFEST).jar

$(MANIFEST).jar: classes $(MANIFEST).mf
	jar cfm $(MANIFEST).jar $(MANIFEST).mf *.class

run: $(MANIFEST).jar
	$(JRUN) -jar $(MANIFEST).jar parameters
