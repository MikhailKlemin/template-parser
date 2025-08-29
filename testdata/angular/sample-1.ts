import { SVGIcon, fileIcon, folderIcon } from "@progress/kendo-svg-icons";
import { CheckBoxState } from "@progress/kendo-angular-inputs";

import { areEqualSampleIterations, ISampleIteration } from "../../models";
import { NlsService } from "@astra/core";

export interface INode {
    icon: SVGIcon;
    text: string;
    isExpanded: boolean;
    checked: CheckBoxState;
    parent?: INode;
    children?: Array<INode>;
    data: ISampleIteration;
}

export function hasChildren(node: INode): boolean {
    return !!node?.children?.length;
}

export function buildNodes(items: Array<ISampleIteration>, nls: NlsService, prevNodes?: Array<INode>): Array<INode> {
    const nodes: Array<INode> = [];
    const itername:string = nls.tr('Iteration') ;
    const samplename:string = nls.tr('Sample') ;

    for (let item of items || []) {
        const existingNode = nodes.find(o => o.data.Index === item.Index);
        const prevNode = findNode(prevNodes, item);
        if (existingNode) {
            if (!existingNode.children) {
                // convert file node to folder node
                existingNode.children = [
                    {
                        parent: existingNode,
                        icon: fileIcon,
                        text:`${itername} ${existingNode.data.Iteration + 1}`,
                        isExpanded: !!prevNode?.isExpanded,
                        checked: prevNode?.checked || false,
                        data: existingNode.data,
                    },
                ];
                existingNode.icon = folderIcon;
                existingNode.data = { ...existingNode.data, Iteration: -1 };
            }
            existingNode.children.push({
                parent: existingNode,
                icon: fileIcon,
                text: `${itername} ${item.Iteration + 1}`,
                isExpanded: !!prevNode?.isExpanded,
                checked: prevNode?.checked || false,
                data: item,
            });
            existingNode.checked = sumChecked(existingNode.children);
        } else {
            const sampleIdStr = item.SampleId ? ` - ${item.SampleId}` : '';
            nodes.push({
                icon: fileIcon,
                text: `${samplename} ${item.Index + 1}${sampleIdStr}`,
                isExpanded: !!prevNode?.isExpanded,
                checked: prevNode?.checked || false,
                data: item,
            });
        }
    }

    return nodes;
}

export function numLevels(nodes: Array<INode>): number {
    if (!nodes?.length) {
        return 0;
    }
    const childLevels = nodes.map(o => numLevels(o.children));
    return 1 + Math.max(...childLevels);
}

export function expandAll(nodes: Array<INode>, expand: boolean = true): void {
    for (let node of nodes || []) {
        if (hasChildren(node)) {
            node.isExpanded = expand;
            expandAll(node.children, expand);
        }
    }
}
export function areAllExpanded(nodes: Array<INode>, isExpanded: boolean = true): boolean | undefined {
    let count = 0;
    for (let node of nodes || []) {
        if (hasChildren(node)) {
            if (node.isExpanded === isExpanded) {
                count++;
            } else {
                return false;
            }
        }
    }
    return count > 0 ? true : undefined;
}

export function checkNode(node: INode, checked: CheckBoxState): void {
    if (!node) {
        return;
    }

    node.checked = checked;
    checkAll(node.children, checked);

    let parent = node.parent;
    while (parent) {
        parent.checked = sumChecked(parent.children);
        parent = parent.parent;
    }
}
export function checkAll(nodes: Array<INode>, checked: CheckBoxState): void {
    for (let node of nodes || []) {
        node.checked = checked;
        if (hasChildren(node)) {
            checkAll(node.children, checked);
        }
    }
}
export function sumChecked(nodes: Array<INode>): CheckBoxState {
    let countChecked = 0;
    let countUnchecked = 0;
    let countIntermediate = 0;
    for (let node of nodes || []) {
        if (node.checked === true) {
            countChecked++;
        } else if (node.checked === false) {
            countUnchecked++;
        } else {
            countIntermediate++;
        }
    }
    if (countIntermediate) {
        return 'indeterminate';
    } else {
        if (countChecked) {
            if (countUnchecked) {
                return 'indeterminate';
            } else {
                return true;
            }
        } else {
            return false;
        }
    }
}

export function addCheckedLeafs(nodes: Array<INode>, checked: CheckBoxState, destination: Array<INode>): void {
    const res: Array<INode> = [];
    for (let node of nodes || []) {
        if (hasChildren(node)) {
            addCheckedLeafs(node.children, checked, destination);
        } else {
            if (node.checked === checked) {
                destination.push(node);
            }
        }
    }
}

export function findCheckedLeafs(nodes: Array<INode>, checked: CheckBoxState): Array<INode> {
    const res: Array<INode> = [];
    addCheckedLeafs(nodes, checked, res);
    return res;
}

export function checkNodes(nodes: Array<INode>, items: Array<ISampleIteration>, checked: CheckBoxState): void {
    for (let item of items || []) {
        const node = findNode(nodes, item);
        checkNode(node, checked);
    }
}

export function findNode(nodes: Array<INode>, item: ISampleIteration): INode {
    for (let node of nodes || []) {
        if (hasChildren(node)) {
            const child = findNode(node.children, item);
            if (child) {
                return child;
            }
        } else {
            if (areEqualSampleIterations(node.data, item)) {
                return node;
            }
        }
    }
    return undefined;
}