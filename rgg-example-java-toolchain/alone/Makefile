
JFLAGS = -g
JC = javac
JAR = jar
JRUN = java
.SUFFIXES: .java .class

.java.class:
	$(JC) $(JFLAGS) $*.java


MAIN = Hello
CLASSES = \
        Foo.java \
        Hello.java

default: classes

classes: $(CLASSES:.java=.class)

clean:
	$(RM) *.class *.jar

run: $(MAIN).class
	$(JRUN) $(MAIN) parameters
