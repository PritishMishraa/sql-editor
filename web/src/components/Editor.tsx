import React, { useEffect, useRef } from "react";
import Editor from "@monaco-editor/react";

interface MonacoEditorProps {
    code: string;
    onChange: (value: string | undefined) => void;
}

const MonacoEditor: React.FC<MonacoEditorProps> = ({ code, onChange }) => {
    const editorRef = useRef<any>(null);

    useEffect(() => {
        if (editorRef.current) {
            editorRef.current.editor.setTheme('light');
            editorRef.current.editor.focus();
        }
    }, []);

    const handleEditorDidMount = (editor: any, _monaco: any) => {
        editorRef.current = editor;
    };

    return (
        <Editor
            className="w-full h-full"
            defaultLanguage="sql"
            defaultValue={code}
            onChange={onChange}
            onMount={handleEditorDidMount}
            options={{
                minimap: { enabled: false },
                wordWrap: 'on',
                scrollBeyondLastLine: false,
                automaticLayout: true,
                fontSize: 16,
                fontFamily: 'Jetbrains Mono',
                renderLineHighlight: 'none',
                overviewRulerLanes: 0,
                glyphMargin: false,
                folding: false,
                lineNumbersMinChars: 3,
            }}
        />
    );
};

export default MonacoEditor;